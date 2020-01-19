package dev.russia9.trainpix.api.parsing;

import dev.russia9.trainpix.api.object.train.Train;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;

public class Parser {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");
    private static int i = 0;

    static Document getPage(String url) throws IOException {
        i++;
        logger.trace(url + " " + i + " " + System.currentTimeMillis());
        return Jsoup
                .connect(url)
                .cookie("lang", "en")
                .get();
    }
}
