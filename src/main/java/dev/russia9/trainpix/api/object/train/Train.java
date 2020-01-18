package dev.russia9.trainpix.api.object.train;

import dev.russia9.trainpix.api.object.infrastructure.Depot;
import dev.russia9.trainpix.api.object.infrastructure.Railway;

import java.util.ArrayList;
import java.util.List;

public class Train {
    /**
     * Trainpix Train ID
     * <p>
     * Example:
     * https://trainpix.org/vehicle/15744/ - 15744
     */
    private int id;

    /**
     * Train Railway
     */
    private Railway railway;

    /**
     * Train Depot
     */
    private Depot depot;

    /**
     * Train model
     */
    private Model model;

    /**
     * Train manufacturer
     * <p>
     * Example:
     * https://trainpix.org/vehicle/130815/ - Ural locomotives
     */
    private String builder = null;

    /**
     * Train identification number
     * <p>
     * Example:
     * https://trainpix.org/vehicle/1375/ - 12236451/12236469
     */
    private String identificationNumber = null;

    /**
     * Train serial type
     * <p>
     * Example:
     * https://trainpix.org/vehicle/9885/ - 62-301
     */
    private String serialType = null;

    /**
     * Train build date
     * <p>
     * Example:
     * https://trainpix.org/vehicle/130815/ - 05.2018
     */
    private String built = null;

    /**
     * Train category
     * <p>
     * Example:
     * https://trainpix.org/vehicle/1375/ - Electric Locomotives
     */
    private String category = "Other";

    /**
     * Train condition
     * <p>
     * Conditions:
     * 1 - In operation
     * 2 - New
     * 3 - Out of order
     * 4 - Written off
     * 5 - Unknown
     * 6 - Transferred to an other depo
     * 7 - KRP/Modernization
     * 8 - Transferred to an other road
     * 9 - Monument
     * 10 - Refurbishment
     */
    private int condition = 5;

    /**
     * Additional information
     */
    private String note = null;

    /**
     * Train photos
     */
    private List<Photo> photoList;

    public Train(int id) {
        this.id = id;
    }

    public Train(int id, boolean photos) {
        this.id = id;
        if(photos) {
            photoList = new ArrayList<>();
        }
    }

    public void setRailway(Railway railway) {
        this.railway = railway;
    }

    public void setDepot(Depot depot) {
        this.depot = depot;
    }

    public void setModel(Model model) {
        this.model = model;
    }

    public void setBuilder(String builder) {
        this.builder = builder;
    }

    public void setIdentificationNumber(String identificationNumber) {
        this.identificationNumber = identificationNumber;
    }

    public void setSerialType(String serialType) {
        this.serialType = serialType;
    }

    public void setBuilt(String built) {
        this.built = built;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public void setCondition(int condition) {
        this.condition = condition;
    }

    public void setNote(String note) {
        this.note = note;
    }

    public void addPhoto(Photo photo) {
        this.photoList.add(photo);
    }
}
