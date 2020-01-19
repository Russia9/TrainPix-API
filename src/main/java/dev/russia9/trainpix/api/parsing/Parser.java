package dev.russia9.trainpix.api.parsing;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.IOException;

public class Parser {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");

    static Document getPage(String url) throws IOException {
        return Jsoup
                .connect(url)
                .cookie("lang", "en")
                .get();
    }
}
