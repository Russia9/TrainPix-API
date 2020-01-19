package dev.russia9.trainpix.api.method.train;

import dev.russia9.trainpix.api.object.train.Train;

import java.util.List;

public class TrainSearch {
    private int error = 0;
    private int size;
    private List<Train> trains;

    public TrainSearch(int error) {
        this.error = error;
    }

    public TrainSearch(int error, List<Train> trains) {
        this.error = error;
        this.size = trains.size();
        this.trains = trains;
    }
}
