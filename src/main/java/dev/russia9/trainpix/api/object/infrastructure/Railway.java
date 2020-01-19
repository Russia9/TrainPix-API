package dev.russia9.trainpix.api.object.infrastructure;

public class Railway {
    /**
     * Trainpix Railway ID
     * <br>
     * Example:
     * https://trainpix.org/railway/2/ - 2
     */
    private int id;

    /**
     * Railway name
     * <p>
     * Example:
     * https://trainpix.org/railway/2/ - Moscow railway
     */
    private String name;

    public Railway(int id, String name) {
        this.id = id;
        this.name = name;
    }
}
