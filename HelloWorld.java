import com.sun.jna.*;
import java.util.*;
import java.lang.Long;

/** Simple example of JNA interface mapping and usage. */
public class HelloWorld {

    // This is the standard, stable way of mapping, which supports extensive
    // customization and mapping of Java to native types.

    public interface CLibrary extends Library {
        void SayHello(long x);
        void ListTopic();
    }

    public static void main(String[] args) {
        CLibrary clib = (CLibrary) Native.load(
            "/home/zhangwm1987/workspace/go-client/client.so", CLibrary.class);

        clib.SayHello(100);
        clib.ListTopic();
    }
}
