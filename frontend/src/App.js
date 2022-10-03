import "./App.css";
import { useEffect, useState } from "react";
import NavBar from "./Components/NavBar";
import { ChakraProvider } from "@chakra-ui/react";
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
import ActivityDescription from "./Pages/ActivityDescription";

import Profile from "./Pages/Profile";

import CreateActivity from "./Pages/CreateActivity";

import CreateItinerary from "./Pages/CreateItinerary";

import CreateReviews from "./Pages/CreateReviews";

import Itineraries from "./Components/ProfileSubPages/Itineraries";
import Bookings from "./Components/ProfileSubPages/Bookings";
import Itinerary from "./Components/Itinerary";

import Activity from "./Components/Activity";

function App() {
  const [imageUrl, setImageUrl] = useState("");

  useEffect(() => {
    setImageUrl(window.sessionStorage.getItem("avatar_file_name"));
  }, []);

  return (
    <ChakraProvider>
      <Router>
        <NavBar imageUrl={imageUrl} setImageUrl={setImageUrl} />
        <Routes>
          <Route path="discover" element={<Discover />} />
          <Route path="reviews" element={<Reviews />} />
          <Route path="signup" element={<SignUp setImageUrl={setImageUrl} />} />
          <Route path="welcome" element={<Welcome />} />
          <Route path="activitydescription" element={<ActivityDescription />} />
          <Route path="profile" element={<Profile imageUrl={imageUrl} />} />
          {/* <Route path="createactivity" element={<CreateActivity/>}/> */}
          <Route path="itineraries" element={<Itineraries />} />
          <Route path="bookings" element={<Bookings />} />
          <Route path="edit-itinerary/:id" element={<Itinerary />} />
          <Route path="activity/:id" element={<Activity />} />
          <Route path="" element={<SignIn setImageUrl={setImageUrl} />} />
          <Route path="createitinerary" element={<CreateItinerary />} />

          <Route path="createreviews" element={<CreateReviews />} />
        </Routes>
      </Router>
    </ChakraProvider>
  );
}

export default App;
