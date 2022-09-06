import './App.css';
import NavBar from './Components/NavBar';
import { ChakraProvider } from '@chakra-ui/react';
import SignIn from './Components/SignIn';


function App() {
  return (
    <ChakraProvider>
      <NavBar/>
      <SignIn/>
    </ChakraProvider>
  );
}

export default App;
