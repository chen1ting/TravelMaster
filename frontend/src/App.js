import logo from './logo.svg';
import './App.css';
import NavBar from './Components/NavBar';
import { ChakraProvider } from '@chakra-ui/react';


function App() {
  return (
    <ChakraProvider>
      <NavBar />
    </ChakraProvider>
  );
}

export default App;
