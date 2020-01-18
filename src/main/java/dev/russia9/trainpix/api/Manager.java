package dev.russia9.trainpix.api;

import dev.russia9.trainpix.api.method.API;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class Manager {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");

    public Manager() {
        new API();
    }
}
