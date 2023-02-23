import { useDisclosure } from '@chakra-ui/react'

import {
  Box,
  Heading,
  Container,
  Text,
  Button,
  Stack,
  createIcon,
} from '@chakra-ui/react';
import RegisterModal from './components/RegisterModal'

export default function Home({handlePetitions, handleSubmit}) {
  const { isOpen, onOpen, onClose } = useDisclosure()

  console.log(isOpen)
  return (
    <>
      <Container maxW={'3xl'}>
        <Stack
          as={Box}
          textAlign={'center'}
          spacing={{ base: 8, md: 14 }}
          py={{ base: 20, md: 36 }}>
          <Heading
            fontWeight={600}
            fontSize={{ base: '2xl', sm: '4xl', md: '6xl' }}
            lineHeight={'110%'}>
            Veterinaria VNET
          </Heading>
          <Text color={'gray.500'}>
            Dentro de tu cuenta podras ver el historico de tus pedidas, registrar tus mascotas y elegir los combos que deseas
          </Text>
          <Stack
            direction={'column'}
            spacing={3}
            align={'center'}
            alignSelf={'center'}
            position={'relative'}>
            <Button
              colorScheme={'green'}
              bg={'green.400'}
              rounded={'full'}
              px={6}
              _hover={{
                bg: 'green.500',
              }}
              onClick={onOpen}
              >
                Registrarse
              {isOpen ? (
                <RegisterModal onClose={onClose} isOpen={isOpen}/>
              ) : ""}
            </Button>
            <Button variant={'link'} colorScheme={'blue'} size={'md'}>
              Logearse
            </Button>
          </Stack>
        </Stack>
      </Container>
    </>
  );
}