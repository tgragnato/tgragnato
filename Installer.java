import java.io.File;
import java.io.IOException;
import java.util.List;
import javafx.application.Application;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.event.ActionEvent;
import javafx.geometry.Insets;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.ListView;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.stage.FileChooser;
import javafx.stage.Stage;
import org.apache.tools.ant.BuildException;

/**
 * The main graphical user interface
 * @author Tommaso Gragnato
 */
public class Installer extends Application {

    @Override
    public void start(Stage primaryStage) throws IOException {

        InstallAPK install = new InstallAPK();

        // Nodes
        ListView<File> table = new ListView<>();
        Button selectADB = new Button();
        Button openFileButton = new Button();
        Button btn = new Button();
        ListView<String> logger = new ListView<>();

        // Select the ADB executable
        selectADB.setText("Select ADB executable");
        selectADB.setOnAction((ActionEvent event) -> {

            FileChooser fileChooser = new FileChooser();
            install.setADB(fileChooser.showOpenDialog(primaryStage));
        });

        // Select APK files for the installation list
        openFileButton.setText("Select APKs");
        openFileButton.setOnAction((ActionEvent event) -> {

            FileChooser fileChooser = new FileChooser();
            List<File> list = fileChooser.showOpenMultipleDialog(primaryStage);
            ObservableList<File> filtered = FXCollections.observableArrayList();
            for (File file : list) {
                if (file.isFile() && file.getPath().toLowerCase().endsWith(".apk")) {
                    filtered.add(file);
                }
            }
            table.setItems(filtered);
        });

        // Install APKs
        btn.setText("Install");
        btn.setOnAction((ActionEvent event) -> {

            for (File file : table.getItems()) {
                install.setAPK(file);
                try {
                    install.execute();
                    ObservableList<String> log = logger.getItems();
                    log.add(file.getPath()+" installed on every device(s)...");
                    logger.setItems(log);
                }
                catch (BuildException ex) {
                    ObservableList<String> log = logger.getItems();
                    log.add(ex.toString());
                    logger.setItems(log);
                }
            }
        });

        // Layouts
        HBox root = new HBox();
        VBox leftColumn = new VBox();
        root.getChildren().addAll(table, leftColumn);
        leftColumn.getChildren().addAll(selectADB, openFileButton, btn, logger);

        // Size and alignment
        leftColumn.setPadding(new Insets(20, 20, 20, 20));
        table.setMinWidth(500);
        selectADB.setMinWidth(460);
        openFileButton.setMinWidth(460);
        btn.setMinWidth(460);
        logger.setMinWidth(460);

        Scene scene = new Scene(root, 1000, 500);

        primaryStage.setTitle("myAPKs");
        primaryStage.setScene(scene);
        primaryStage.setResizable(false);
        primaryStage.show();
    }

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        launch(args);
    }

}