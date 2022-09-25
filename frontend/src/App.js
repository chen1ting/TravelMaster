import "./App.css";
import NavBar from "./Components/NavBar";
import {ChakraProvider} from "@chakra-ui/react";
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link as RouteLink,
} from "react-router-dom";
import SignIn from "./Components/SignIn";
import Welcome from "./Components/Welcome";
import SignUp from "./Components/SignUp";
import Discover from "./Pages/Discover";
import Reviews from "./Pages/Reviews";
import ActivityDescription from './Pages/ActivityDescription';
import Profile from "./Pages/Profile";
import CreateEvent from './Pages/CreateEvent';
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
                    <Route path="createevent" element={<CreateEvent/>}
                    />
                </Routes>
            </Router>
        </ChakraProvider>
    );
}

export default App;
