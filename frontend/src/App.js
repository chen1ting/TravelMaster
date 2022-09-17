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
import Discover from "./Components/Discover";
import Reviews from "./Components/Reviews";


function App() {
    return (
        <ChakraProvider>
            <Router>
                <NavBar/>
                <Routes>
                    <Route
                        path="discover"
                        element={<Discover/>}
                    />
                    <Route
                        path="reviews"
                        element={<Reviews/>}
                    />
                    <Route
                        path="signup"
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
