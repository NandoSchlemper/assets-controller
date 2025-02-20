import { createStaticNavigation } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { Text, View } from "react-native";


function HomeScreen() {
    // Definimos um estilo
    // Algum conteúdo para o View e é isso
    return(
        <View style={{flex:1, alignItems:'center', justifyContent:'center'}}>
            <Text>Home bruh</Text>
        </View>
    )
}

// Aparentemente, esse rootstack me permite adicionar todas as telas que vão
// Ser usadas na parte da navegação
const RootStack = createNativeStackNavigator({
    screens: {
        Home: HomeScreen
    }
})

// Aqui é a instancia da navegação estática baseado no Stack que criamos.
const Navigation = createStaticNavigation(RootStack)

export default function App() {
    return <Navigation />
}
