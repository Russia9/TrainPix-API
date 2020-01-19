package dev.russia9.trainpix.api.object.train;

public class Model {
    /**
     * Trainpix Model ID
     * <br>
     * Example:
     * https://trainpix.org/list.php?mid=98 - 98
     */
    private int id;

    /**
     * Model name
     *
     * Example:
     * https://trainpix.org/list.php?mid=98 - EM2
     */
    private String name;

    public Model(int id, String name) {
        this.id = id;
        this.name = name;
    }
}
