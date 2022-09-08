import './App.css';
import NavBar from './Components/NavBar';
import { ChakraProvider } from '@chakra-ui/react';
import SignIn from './Components/SignIn';
import SignUp from "./Components/SignUp";


function App() {
  return (
    <ChakraProvider>
      <NavBar/>
      <SignIn/>
    </ChakraProvider>
  );
}

export default App;
