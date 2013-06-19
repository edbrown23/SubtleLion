/**
 * This code falls under the MIT license, found in full in the file license.txt
 * Created By: Eric Brown
 * Date: 6/18/2013
 */
import com.badlogic.gdx.backends.lwjgl.LwjglApplication;

public class Main {
    public static void main(String[] args){
        new LwjglApplication(new SubtleLion(), "SubtleLion", 320, 240, false);
    }
}
