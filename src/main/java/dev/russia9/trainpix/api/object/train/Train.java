package dev.russia9.trainpix.api.object.train;

import dev.russia9.trainpix.api.object.infrastructure.Depot;
import dev.russia9.trainpix.api.object.infrastructure.Railway;

import java.util.ArrayList;
import java.util.List;

public class Train {
    /**
     * Trainpix TrainAPI ID
     * <p>
     * Example:
     * https://trainpix.org/vehicle/15744/ - 15744
     */
    private int id;


    /**
     * TrainAPI name
     * <p>
     * Example:
     * https://trainpix.org/vehicle/15744/ - ЭМ4-015
     */
    private String name = null;

    /**
     * TrainAPI Railway
     */
    private Railway railway;

    /**
     * TrainAPI Depot
     */
    private Depot depot;

    /**
     * TrainAPI model
     */
    private Model model;

    /**
     * TrainAPI manufacturer
     * <p>
     * Example:
     * https://trainpix.org/vehicle/130815/ - Ural locomotives
     */
    private String builder = null;

    /**
     * TrainAPI identification number
     * <p>
     * Example:
     * https://trainpix.org/vehicle/1375/ - 12236451/12236469
     */
    private String identificationNumber = null;

    /**
     * TrainAPI serial type
     * <p>
     * Example:
     * https://trainpix.org/vehicle/9885/ - 62-301
     */
    private String serialType = null;

    /**
     * TrainAPI build date
     * <p>
     * Example:
     * https://trainpix.org/vehicle/130815/ - 05.2018
     */
    private String built = null;

    /**
     * TrainAPI category
     * <p>
     * Example:
     * https://trainpix.org/vehicle/1375/ - Electric Locomotives
     */
    private String category = "Other";

    /**
     * TrainAPI condition
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
     * TrainAPI photos
     */
    private List<Photo> photoList;

    public Train(int id) {
        this.id = id;
    }

    public Train(int id, boolean photos) {
        this.id = id;
        if (photos) {
            photoList = new ArrayList<>();
        }
    }

    public void addPhoto(Photo photo) {
        this.photoList.add(photo);
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Railway getRailway() {
        return railway;
    }

    public void setRailway(Railway railway) {
        this.railway = railway;
    }

    public Depot getDepot() {
        return depot;
    }

    public void setDepot(Depot depot) {
        this.depot = depot;
    }

    public Model getModel() {
        return model;
    }

    public void setModel(Model model) {
        this.model = model;
    }

    public String getBuilder() {
        return builder;
    }

    public void setBuilder(String builder) {
        this.builder = builder;
    }

    public String getIdentificationNumber() {
        return identificationNumber;
    }

    public void setIdentificationNumber(String identificationNumber) {
        this.identificationNumber = identificationNumber;
    }

    public String getSerialType() {
        return serialType;
    }

    public void setSerialType(String serialType) {
        this.serialType = serialType;
    }

    public String getBuilt() {
        return built;
    }

    public void setBuilt(String built) {
        this.built = built;
    }

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public int getCondition() {
        return condition;
    }

    public void setCondition(int condition) {
        this.condition = condition;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }
}
