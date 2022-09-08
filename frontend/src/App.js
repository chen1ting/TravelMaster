import './App.css';
import NavBar from './Components/NavBar';
import { ChakraProvider } from '@chakra-ui/react';
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link as RouteLink, BrowserRouter
} from "react-router-dom";
import SignIn from './Components/SignIn';
import SignUp from "./Components/SignUp";


function App() {
  return (
    <ChakraProvider>

        <BrowserRouter>
            <NavBar/>
            <Routes>
                <Route
                    path="SignUp"
                    element={<SignUp />}
                />
                <Route
                    path=""
                    element={<SignIn />}
                />
            </Routes>
        </BrowserRouter>


    </ChakraProvider>
  );
}

export default App;
