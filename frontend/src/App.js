import './App.css';
import NavBar from './Components/NavBar';
import {ChakraProvider} from '@chakra-ui/react';
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link as RouteLink
} from "react-router-dom";
import SignIn from './Components/SignIn';
import SignUp from "./Components/SignUp";


function App() {
    return (
        <ChakraProvider>
            <Router>
                <NavBar/>
                <Routes>
                    <Route
                        path="SignUp"
                        element={<SignUp/>}
                    />
                    <Route
                        path=""
                        element={<SignIn/>}
                    />
                </Routes>
            </Router>
        </ChakraProvider>
    );
}

export default App;
