public class LogLevels {
    
    public static String message(String logLine) {
        String message = "";
        if (logLine.startsWith("[INFO]")) {
            message = logLine.substring(7);
        }
        if (logLine.startsWith("[ERROR]")) {
            message = logLine.substring(8);
        }
        if (logLine.startsWith("[WARNING]")) {
            message = logLine.substring(10);
        }
        return message.strip();
    }

    public static String logLevel(String logLine) {
        String res = "";
        if (logLine.startsWith("[INFO]")) {
            res = "info";
        }
        if (logLine.startsWith("[ERROR]")) {
            res = "error";
        }
        if (logLine.startsWith("[WARNING]")) {
            res = "warning";
        }
        return res;
    }

    public static String reformat(String logLine) {
        String message = message(logLine);
        String level = logLevel(logLine);
        return String.format("%s (%s)", message, level);
    }
}
