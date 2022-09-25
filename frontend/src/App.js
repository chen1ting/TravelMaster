
import "./App.css";
import NavBar from "./Components/NavBar";
import {ChakraProvider} from "@chakra-ui/react";

import {
    BrowserRouter as Router,
    Routes,
    Route,

    Link as RouteLink,

} from "react-router-dom";
import SignIn from './Components/SignIn';
import SignUp from "./Components/SignUp";
import Discover from "./Pages/Discover";
import Reviews from "./Pages/Reviews";
import ActivityDescription from './Pages/ActivityDescription';

import CreateEvent from './Pages/CreateEvent';
import Profile from "./Pages/Profile";
import Welcome from "./Pages/Welcome";



function App() {
    return (
        <ChakraProvider>
            <Router>
                <NavBar/>
                <Routes>

                    <Route path="discover" element={<Discover/>}/>
                    <Route path="reviews" element={<Reviews/>}/>
                    <Route path="signup" element={<SignUp/>}/>
                    <Route path="welcome" element={<Welcome/>}/>
                    <Route path="activitydescription" element={<ActivityDescription/>}/>
                    <Route path="profile" element={<Profile/>}/>
                    <Route path="" element={<SignIn/>}/>

                </Routes>
            </Router>
        </ChakraProvider>
    );
}

export default App;
