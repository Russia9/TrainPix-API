package dev.russia9.trainpix.api.method;

import com.google.gson.Gson;
import dev.russia9.trainpix.api.parsing.TrainParser;

import static spark.Spark.*;

public class API {
    public API() {
        Gson gson = new Gson();
        port(8080);
        path("/api", () -> {
            path("/train", () -> {
                get("/search", (request, response) -> {
                    response.type("application/json");
                    String query = request.queryParamOrDefault("term", "ЭР2");
                    int size = Integer.parseInt(request.queryParamOrDefault("size", "5"));
                    if(size>20) size = 20;
                    return gson.toJson(TrainParser.search(query, size, null));
                });
            });
        });
        internalServerError((request, response) -> {
            response.type("application/json");
            return "{}";
        });
    }
}
