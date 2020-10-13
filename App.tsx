import * as eva from "@eva-design/eva";
import {
  ApplicationProvider,
  Button,
  Layout,
  Text,
  TopNavigation,
} from "@ui-kitten/components";
import { Audio } from "expo-av";
import React from "react";
import { Alert, StatusBar, useColorScheme } from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";

const HomeScreen = () => {
  const recording = React.useRef<Audio.Recording>();
  const [isRecording, setIsRecording] = React.useState(false);

  React.useEffect(() => {
    (async () => {
      const permissionResponse = await Audio.requestPermissionsAsync();
      if (!permissionResponse.granted) {
        Alert.alert(
          "Permission error!",
          "You must enable audio permissions!",
          [{ text: "OK", onPress: () => console.log("OK Pressed") }],
          { cancelable: false }
        );
      }
    })();
  }, []);

  return (
    <>
      <TopNavigation title="Nectar" />
      <Layout
        style={{
          flex: 1,
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Button
          appearance="outline"
          onPress={async () => {
            if (!isRecording) {
              recording.current = new Audio.Recording();
              await recording.current.prepareToRecordAsync(
                Audio.RECORDING_OPTIONS_PRESET_HIGH_QUALITY
              );
              await recording.current.startAsync();
              setIsRecording(true);
            } else {
              await recording.current!.stopAndUnloadAsync();
              setIsRecording(false);
            }
          }}
        >
          Press to toggle Recording!
        </Button>
        <Text>{isRecording ? "Recording..." : "Not Recording..."}</Text>
      </Layout>
    </>
  );
};

export default function App() {
  const colorScheme = useColorScheme();

  return (
    <ApplicationProvider {...eva} theme={eva[colorScheme!]}>
      <SafeAreaProvider>
        <HomeScreen />
      </SafeAreaProvider>
      <StatusBar translucent={false} />
    </ApplicationProvider>
  );
}
