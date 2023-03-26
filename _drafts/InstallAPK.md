---
title: InstallAPK
layout: default
---

```java
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
```

```java
import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.PrintStream;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import org.apache.tools.ant.BuildException;
import org.apache.tools.ant.Task;

/**
 * Custom task that uses the ADB tool to install a specified APK on all
 * connected devices and emulators.
 * @author Daniel Dyer
 */
public class InstallAPK extends Task
{
    private File apkFile;

    public void setAPK(File apkFile)
    {
        this.apkFile = apkFile;
    }

    @Override
    public void execute() throws BuildException
    {
        if (apkFile == null)
        {
            throw new BuildException("APK file must be specified");
        }
        try
        {
            List<String> devices = getDeviceIdentifiers();
            System.out.printf("Installing %s on %d device(s)...%n", apkFile, devices.size());
            ExecutorService executor = Executors.newFixedThreadPool(devices.size());
            List<Future<Void>> futures = new ArrayList<Future<Void>>(devices.size());
            for (final String device : devices)
            {
                futures.add(executor.submit(new Callable<Void>()
                {
                    public Void call() throws IOException, InterruptedException
                    {
                        installOnDevice(device);
                        return null;
                    }
                }));
            }
            for (Future<Void> future : futures)
            {
                future.get();
            }
            executor.shutdown();
            executor.awaitTermination(60, TimeUnit.SECONDS);
        }
        catch (Exception ex)
        {
            throw new BuildException(ex);
        }
    }

    private void installOnDevice(String device) throws IOException, InterruptedException
    {
        String[] command = new String[]{"adb", "-s", device, "install", "-r", apkFile.toString()};
        Process process = Runtime.getRuntime().exec(command);
        consumeStream(process.getInputStream(), System.out, device);
        if (process.waitFor() != 0)
        {
            consumeStream(process.getErrorStream(), System.err, device);
            throw new BuildException(String.format("Installing APK on %s failed.", device));
        }
    }

    private void consumeStream(InputStream in, PrintStream out, String tag) throws IOException
    {
        BufferedReader reader = new BufferedReader(new InputStreamReader(in));
        try
        {
            for (String line = reader.readLine(); line != null; line = reader.readLine())
            {
                out.println(tag != null ? String.format("[%s] %s", tag, line.trim()) : line);
            }
        }
        finally
        {
            reader.close();
        }
    }

    private List<String> getDeviceIdentifiers() throws IOException, InterruptedException
    {
        Process process = Runtime.getRuntime().exec("adb devices");
        List devices = new ArrayList<String>(10);
        BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
        try
        {
            for (String line = reader.readLine(); line != null; line = reader.readLine())
            {
                if (line.endsWith("device"))
                {
                    devices.add(line.split("s")[0]);
                }
            }
            if (process.waitFor() != 0)
            {
                consumeStream(process.getErrorStream(), System.err, null);
                throw new BuildException("Failed getting list of connected devices/emulators.");
            }
        }
        finally
        {
            reader.close();
        }
        return devices;
    }
}
```
