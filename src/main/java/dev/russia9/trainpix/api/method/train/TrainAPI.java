package dev.russia9.trainpix.api.method.train;

import com.google.gson.Gson;
import dev.russia9.trainpix.api.exception.TrainNotFoundException;
import dev.russia9.trainpix.api.parsing.TrainParser;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import spark.Request;
import spark.Response;

import java.io.IOException;
import java.util.List;

public class TrainAPI {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");
    private static Gson gson = new Gson();

    public static String search(Request request, Response response) {
        response.type("application/json");
        String query = request.queryParamOrDefault("query", "ЭР2");
        int size = Integer.parseInt(request.queryParamOrDefault("size", "5"));
        logger.debug("Search: query:'" + query + "' size:" + size);
        if (size > 20) size = 20;
        TrainSearch result;
        try {
            List<dev.russia9.trainpix.api.object.train.Train> trainList = TrainParser.search(query, size, false, null);
            result = new TrainSearch(0, trainList);
        } catch (TrainNotFoundException | IOException e) {
            result = new TrainSearch(404);
        }
        return gson.toJson(result);
    }

    public static String quickSearch(Request request, Response response) {
        response.type("application/json");
        String query = request.queryParamOrDefault("query", "ЭР2");
        int size = Integer.parseInt(request.queryParamOrDefault("size", "5"));
        logger.debug("QSearch: query:'" + query + "' size:" + size);
        TrainSearch result;
        try {
            List<dev.russia9.trainpix.api.object.train.Train> trainList = TrainParser.search(query, size, true, null);
            result = new TrainSearch(0, trainList);
        } catch (TrainNotFoundException | IOException e) {
            result = new TrainSearch(404);
        }
        return gson.toJson(result);
    }
}
