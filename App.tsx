import * as eva from '@eva-design/eva';
import { ApplicationProvider, Layout, Text, TopNavigation } from '@ui-kitten/components';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { useColorScheme } from 'react-native';
import { SafeAreaProvider } from 'react-native-safe-area-context';

const HomeScreen = () => (
  <Layout style={{ flex: 1, justifyContent: 'center', alignItems: 'center', width: '100%', height: '100%' }}>
    <TopNavigation
      title='Nectar'
    />
    <Text category='h1'>Hi!</Text>
  </Layout>
);

export default function App() {
  const colorScheme = useColorScheme()

  return (
    <ApplicationProvider {...eva} theme={eva[colorScheme!]}>
      <SafeAreaProvider>
        <HomeScreen />
        <StatusBar />
      </SafeAreaProvider>
    </ApplicationProvider>
  )
};
