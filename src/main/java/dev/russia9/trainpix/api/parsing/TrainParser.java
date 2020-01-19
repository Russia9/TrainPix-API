package dev.russia9.trainpix.api.parsing;

import dev.russia9.trainpix.api.exception.TrainNotFoundException;
import dev.russia9.trainpix.api.object.infrastructure.Depot;
import dev.russia9.trainpix.api.object.infrastructure.Railway;
import dev.russia9.trainpix.api.object.train.Model;
import dev.russia9.trainpix.api.object.train.Train;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

import java.io.IOException;
import java.net.URLEncoder;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import static dev.russia9.trainpix.api.parsing.Parser.getPage;

public class TrainParser {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");

    public static List<Train> search(String query, int size, boolean quick, Map<String, Integer> params) throws TrainNotFoundException, IOException {
        List<Train> result = new ArrayList<>();

        StringBuilder searchURI = new StringBuilder();
        searchURI.append("https://trainpix.org/vsearch.php");
        searchURI.append("?num=").append(URLEncoder.encode(query));

        Document searchPage = getPage(searchURI.toString());
        if (searchPage.getElementsContainingOwnText("Nothing found.").size() > 0) {
            throw new TrainNotFoundException();
        }

        Element searchTable = searchPage.getElementsByClass("rtable").first().getElementsByTag("tbody").first();
        Elements trainRows = searchTable.getElementsByAttributeValueStarting("class", "s");

        int i = 0;
        for (Element trainRow : trainRows) {
            if (i >= size) {
                break;
            }
            int id = Integer.parseInt(trainRow.getElementsByTag("a").first().attr("href").split("/")[2]);

            int condition = Integer.parseInt(trainRow.className().substring(1));
            if (condition > 10) condition -= 10;

            if (condition == 6 || condition == 8) {
                continue;
            }

            Train train;

            if (quick) {
                train = new Train(id);
                train.setCondition(condition);
                train.setName(trainRow.getElementsByTag("a").first().ownText());
            } else {
                train = get(id);
            }

            result.add(train);
            i++;
        }
        return result;
    }

    public static Train get(int id) throws TrainNotFoundException, IOException {
        Train train = new Train(id);
        String trainURI = "https://trainpix.org/vehicle/" + id + "/";
        Document trainPage = getPage(trainURI);

        if (trainPage.getElementsContainingOwnText("The rail vehicle is not found").size() > 0) {
            throw new TrainNotFoundException();
        }

        Elements infoRows = trainPage.getElementsByClass("h21");

        String name = trainPage.getElementsByTag("h1").first().ownText();
        train.setName(name);
        for (Element infoRow : infoRows) {
            if (infoRow.children().size() > 1) {
                if (infoRow.getElementsContainingOwnText("Railway District/Company:").size() > 0) {
                    int railwayId = Integer.parseInt(infoRow.getElementsByTag("a").first().attr("href").split("/")[2]);
                    String railwayName = infoRow.getElementsByTag("a").first().ownText();
                    train.setRailway(new Railway(railwayId, railwayName));
                } else if (infoRow.getElementsContainingOwnText("Depot:").size() > 0) {
                    int depotId = Integer.parseInt(infoRow.getElementsByTag("a").first().attr("href").split("=")[1]);
                    String depotName = infoRow.getElementsByTag("a").first().ownText();
                    train.setDepot(new Depot(depotId, depotName));
                } else if (infoRow.getElementsContainingOwnText("Model:").size() > 0) {
                    int modelId = Integer.parseInt(infoRow.getElementsByTag("a").first().attr("href").split("=")[1]);
                    String modelName = infoRow.getElementsByTag("a").first().ownText();
                    train.setModel(new Model(modelId, modelName));
                } else if (infoRow.getElementsContainingOwnText("Builder:").size() > 0) {
                    train.setBuilder(infoRow.getElementsByClass("d").first().ownText());
                } else if (infoRow.getElementsContainingOwnText("Identification number:").size() > 0) {
                    train.setIdentificationNumber(infoRow.getElementsByTag("b").first().ownText());
                } else if (infoRow.getElementsContainingOwnText("Serial type:").size() > 0) {
                    train.setSerialType(infoRow.getElementsByTag("b").first().ownText());
                } else if (infoRow.getElementsContainingOwnText("Built:").size() > 0) {
                    train.setBuilt(infoRow.getElementsByTag("b").first().ownText());
                } else if (infoRow.getElementsContainingOwnText("Category:").size() > 0) {
                    train.setCategory(infoRow.getElementsByClass("d").first().ownText());
                } else if (infoRow.getElementsContainingOwnText("Note:").size() > 0) {
                    train.setNote(infoRow.getElementsByClass("d").first().ownText());
                }
            }
        }

        return train;
    }
}
