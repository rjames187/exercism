public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        int raw = speed * 221;
        if (speed >= 5 && speed <= 8) {
            return raw * 0.9;
        }
        if (speed == 9) {
            return raw * 0.8;
        }
        if (speed == 10) {
            return raw * 0.77;
        }
        return (double) raw;
    }

    public int workingItemsPerMinute(int speed) {
        return (int) productionRatePerHour(speed) / 60;
    }
}
