package dev.russia9.trainpix.api.method;

import com.google.gson.Gson;
import dev.russia9.trainpix.api.parsing.Parser;

import static spark.Spark.*;

public class API {
    public API() {
        Gson gson = new Gson();
        port(8080);
        path("/api", () -> {
            path("/train", () -> {
                get("/search", (request, response) -> {
                    response.type("application/json");
                    return gson.toJson(Parser.searchTrains("dfd",1));
                });
            });
        });
    }
}
