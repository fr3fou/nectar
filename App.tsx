import * as eva from "@eva-design/eva";
import {
  ApplicationProvider,
  Layout,
  Text,
  TopNavigation,
} from "@ui-kitten/components";
import React from "react";
import { StatusBar, useColorScheme } from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";

const HomeScreen = () => (
  <>
    <TopNavigation title="Nectar" />
    <Layout
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text>h</Text>
    </Layout>
  </>
);

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
