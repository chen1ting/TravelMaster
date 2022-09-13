<<<<<<< Updated upstream
import { useState } from 'react';
import { Flex, Box, Text, useColorModeValue, Link, Menu, MenuButton, Button, Avatar, MenuList} from '@chakra-ui/react';
=======
import { 
  Flex, 
  Box, 
  useColorModeValue, 
  Link, 
  Menu, 
  MenuButton, 
  Button, 
  Avatar, 
  MenuList
} from '@chakra-ui/react';
>>>>>>> Stashed changes

const MenuItem = ({ children, to = '/'}) => {
  return (
    <Link 
      px = "2"
      py = "1"
      rounded = "md"
      _hover = {{
        bg: useColorModeValue('blue.100', 'blue.700'),
        fontWeight: "bold",
      }}
      href={to}>
      {children}
    </Link>
  );
};

const Header = (props) => {
<<<<<<< Updated upstream
    const [show, setShow] = useState(false);
    return (
        <Flex
            // mb={4}
            p={2}
            as="nav"
            align="center"
            justify="space-between"
            wrap="wrap"
            w="100%"
            bg="cornflowerblue"
        >
          <Box w="200px" p="-1">
              <Text as="i" fontSize="5xl" fontWeight="bold">
                  TM
              </Text>
          </Box>

            <Box
                display={{base: show ? 'block' : 'none', md: 'block'}}
                flexBasis={{base: '100%', md: 'auto'}}
            >
                <Flex
                >
                    <MenuItem to="/discover">Discover</MenuItem>
                    <MenuItem to="/createitinery">Create an Itinerary</MenuItem>
                    <MenuItem to="/itineraries">Itineraries</MenuItem>
                    <MenuItem to="/bookings">Bookings</MenuItem>
                    <Flex alignItems={'center'}>
                        <Menu>
                            <MenuButton
                                as={Button}
                                rounded={'full'}
                                variant={'link'}
                                cursor={'pointer'}
                                minW={0}>
                                <Avatar
                                    size={'sm'}
                                    src={
                                        'https://www.cleanpng.com/png-person-silhouette-download-human-vector-1709279/://images.unsplash.com/photo-1493666438817-866a91353ca9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9'
                                    }
                                />
                            </MenuButton>
                            <MenuList justify="center">
                                <MenuItem>Profile</MenuItem>
                                <MenuItem>Log Out</MenuItem>
                            </MenuList>
                        </Menu>
                    </Flex>
                </Flex>
            </Box>
        </Flex>
    );
=======
  return (
    <Flex
      mb={4}
      p={4}
      as="nav"
      justify="space-between"
      bg="blue.500"
    > 
      <Box>Travel Master Logo</Box>
      <Flex>
        <MenuItem mr = "2" to="/createItinerary">Create an Itinerary</MenuItem>
        <MenuItem mr = "2" to="/itineraries">Itineraries</MenuItem>
        <MenuItem mr = "2" to="/bookings">Bookings</MenuItem>

        <Menu w = "50%">
          <MenuButton as={Button} rounded={'full'} variant={'link'}>
            <Avatar
              ml = "2"
              mr = "2"
              size={'sm'}
              src={
                'https://www.cleanpng.com/png-person-silhouette-download-human-vector-1709279/://images.unsplash.com/photo-1493666438817-866a91353ca9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9'
              }
            />
          </MenuButton>

          <MenuList size = "50%">
            <Flex direction = "column">
              <MenuItem>Log In</MenuItem>
              <MenuItem to="/SignUp">Sign Up</MenuItem>
            </Flex>
          </MenuList>

        </Menu>
      </Flex>
    </Flex>
  );
>>>>>>> Stashed changes
};

export default Header;
