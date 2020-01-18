package dev.russia9.trainpix.api.parsing;

import dev.russia9.trainpix.api.object.train.Train;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Parser {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");

    public static Train getTrain(int id) {
        Train train = new Train(id);

        train.setNote("Test note");
        return train;
    }

    public static List<Train> searchTrains(String query, int size) {
        List<Train> result = new ArrayList<>();
        Train train = new Train(123123);
        train.setBuilder("Test builder");
        train.setNote("Test note");
        train.setCondition(1);
        result.add(train);
        result.add(train);
        result.add(train);
        return result;
    }

    public static Document getPage(String url, String lang) throws IOException {
        return Jsoup
                .connect(url)
                .cookie("lang", lang)
                .get();
    }
}
