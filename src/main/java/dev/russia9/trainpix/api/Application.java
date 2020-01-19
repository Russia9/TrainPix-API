package dev.russia9.trainpix.api;

import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.core.config.Configurator;

/**
 * Launch class
 *
 * @author Russia9
 * @since v0.0.1
 */
public class Application {
    private static final Logger logger = LogManager.getLogger("TrainPixAPI");

    public static void main(String[] args) {
        System.setProperty(org.slf4j.impl.SimpleLogger.DEFAULT_LOG_LEVEL_KEY, "Off");
        if (System.getenv("LEVEL") != null) {
            String level = System.getenv("LEVEL");
            switch (level) {
                case "OFF":
                    Configurator.setLevel("TrainPixAPI", Level.OFF);
                    break;
                case "FATAL":
                    Configurator.setLevel("TrainPixAPI", Level.FATAL);
                    break;
                case "ERROR":
                    Configurator.setLevel("TrainPixAPI", Level.ERROR);
                    break;
                case "WARN":
                    Configurator.setLevel("TrainPixAPI", Level.WARN);
                    break;
                case "DEBUG":
                    Configurator.setLevel("TrainPixAPI", Level.DEBUG);
                    break;
                case "TRACE":
                    Configurator.setLevel("TrainPixAPI", Level.TRACE);
                    break;
                default:
                    Configurator.setLevel("TrainPixAPI", Level.INFO);
                    break;
            }
        }
        for (String arg : args) {
            if ("Debug=true".equals(arg)) {
                Configurator.setLevel("TrainPixAPI", Level.DEBUG);
            } else if ("Trace=true".equals(arg)) {
                Configurator.setLevel("TrainPixAPI", Level.TRACE);
            }
        }

        logger.info("TrainPix API by Russia9");
        logger.info("For more information: https://github.com/Russia9/TrainPix-API");

        new Manager();
    }
}
