package dev.russia9.trainpix.api.method;

import com.google.gson.Gson;
import dev.russia9.trainpix.api.method.train.TrainAPI;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import static spark.Spark.*;

public class API {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");

    public API() {
        Gson gson = new Gson();
        port(10000);
        path("/api", () -> {
            path("/train", () -> {
                get("/search", TrainAPI::search);
                get("/qsearch", TrainAPI::quickSearch);
            });
        });
        internalServerError((request, response) -> {
            response.type("application/json");
            return "{}";
        });
    }
}
