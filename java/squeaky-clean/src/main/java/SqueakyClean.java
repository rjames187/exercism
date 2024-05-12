class SqueakyClean {
    static String clean(String identifier) {
        identifier = identifier.replace(' ', '_').replace('7', 't').replace('1', 'l');
        identifier = identifier.replace('0', 'o').replace('3', 'e').replace('4', 'a');
        StringBuilder str = new StringBuilder();
        str.append(identifier);
        int dash_idx = identifier.indexOf('-');
        if (dash_idx > -1) {
            char toCapitalize = str.charAt(dash_idx + 1);
            str.setCharAt(dash_idx + 1, Character.toUpperCase(toCapitalize));
            str.deleteCharAt(dash_idx);
        }
        identifier = str.toString();
        return identifier.replaceAll("[!$#.¡]+", "");
    }
}
