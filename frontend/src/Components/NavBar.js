import {useState} from 'react';
import {
    Flex,
    Box,
    Text,
    useColorModeValue,
    Menu,
    MenuButton,
    Button,
    Avatar,
    MenuList,
    MenuItem,
} from '@chakra-ui/react';
import {useNavigate} from "react-router-dom";

const paddingSpace = '20px'

const NavButton = (displayString, navigateString) => {
    const navigate = useNavigate();
    return (
        <MenuButton
            px="2"
            py="1"
            rounded='md'
            onClick={() => {
                navigate(navigateString);
            }}
            _hover={{
                bg: useColorModeValue('blue.200', 'blue.400'),
                fontWeight: "bold",
            }}
            m={2}
        >
            {displayString}
        </MenuButton>
    )
}

const Header = () => {
    const navigate = useNavigate();
    const [show] = useState(false);

    return (
        <Flex
            // mb={4}
            p={2}
            as="nav"
            alignItems="center"
            justify="space-between"
            w="100%"
            bg="cornflowerblue"
            h="6vh"
        >
            <Box w="200px" p="-1" paddingLeft={paddingSpace}>
                <Text as="i" fontSize="5xl" fontWeight="bold">
                    TM
                </Text>
            </Box>

            <Box
                display={{base: show ? 'block' : 'none', md: 'block'}}
                flexBasis={{base: '100%', md: 'auto'}}
                paddingRight={paddingSpace}
            >

                <Flex fontSize={'xl'} px="10">
                    {/* Navbar buttons */}
                    <Menu>
                        {NavButton('Discover', '/discover')}
                        {NavButton('Create an Itinerary', '/createitinery')}
                        {NavButton('Itineraries', '/itineraries')}
                        {NavButton('Bookings', '/bookings')}
                    </Menu>

                    {/* User Icon (Profile) */}

                    <Flex alignItems={'center'} px="3">
                        <Menu>
                            <MenuButton
                                as={Button}
                                rounded={'full'}
                                variant={'link'}
                                cursor={'pointer'}
                                minW={0}
                                paddingLeft={'14px'}
                            >
                                {/* Replace src with user image url, with the default as current image*/}
                                <Avatar
                                    size={'md'}
                                    src={
                                        'https://www.cleanpng.com/png-person-silhouette-download-human-vector-1709279/://images.unsplash.com/photo-1493666438817-866a91353ca9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9'
                                    }
                                />
                            </MenuButton>
                            <MenuList justify="center">
                                <MenuItem
                                    onClick={() => {
                                        navigate('/profile');
                                    }}
                                    px="2"
                                    py="1"
                                    rounded='md'
                                    _hover={{
                                        bg: useColorModeValue('blue.200', 'blue.400'),
                                        fontWeight: "bold",
                                    }}
                                >
                                    Profile
                                </MenuItem>
                                <MenuItem
                                    onClick={() => {
                                        navigate('/');
                                    }}
                                    px="2"
                                    py="1"
                                    rounded='md'
                                    _hover={{
                                        bg: useColorModeValue('blue.200', 'blue.400'),
                                        fontWeight: "bold",
                                    }}
                                >
                                    Log Out
                                </MenuItem>
                            </MenuList>
                        </Menu>
                    </Flex>
                </Flex>
            </Box>
        </Flex>
    );


};

export default Header;
