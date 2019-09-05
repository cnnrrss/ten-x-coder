import java.util.ArrayList;

public class Generics {

    public static void main(String[] args) {

        List v = new ArrayList();
        v.add("test"); // A String that cannot be cast to an Integer
        Integer i = (Integer)v.get(0); // Run time error
        
        List<String> w = new ArrayList<String>();
        w.add("test");
        Integer j = v.get(0); // (type error), compilation-time error

    }
}