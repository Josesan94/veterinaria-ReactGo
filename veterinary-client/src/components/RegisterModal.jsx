import React from 'react'
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalFooter,
    ModalBody,
    Button,
  } from '@chakra-ui/react'
  import { useFormik } from "formik";
import {
    Box,
    Checkbox,
    Flex,
    FormControl,
    FormLabel,
    Input,
    VStack,
    Select
  } from "@chakra-ui/react";

const RegisterModal = ({onClose, isOpen}) => {

    const formik = useFormik({
        initialValues: {
          email: "",
          name:"",
          password: "",
          pet: ""
        },
        onSubmit: (values) => {
          alert(JSON.stringify(values, null, 2))
        },
      })


  return (
    <Modal  isCentered isOpen >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader align="center">Registro</ModalHeader>
          <ModalBody>
          <Flex bg="gray.100" align="center" justify="center">
      <Box bg="white" p={6} rounded="md">
        <form onSubmit={formik.handleSubmit}>
          <VStack spacing={4} align="flex-start">
            <FormControl>
              <FormLabel htmlFor="email">Email</FormLabel>
              <Input
                id="email"
                name="email"
                type="email"
                variant="filled"
                onChange={formik.handleChange}
                value={formik.values.email}
              />
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="name">Nombre</FormLabel>
              <Input
                id="name"
                name="name"
                type="name"
                variant="filled"
                onChange={formik.handleChange}
                value={formik.values.name}
              />
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="password">Password</FormLabel>
              <Input
                id="password"
                name="password"
                type="password"
                variant="filled"
                onChange={formik.handleChange}
                value={formik.values.password}
              />
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="pet">Elija su mascota</FormLabel>
              <Select
                id="pet"
                name="pet"
                type="pet"
                variant="filled"
                onChange={formik.handleChange}
                value={formik.values.pet}
              >
                <option value='cat'>Gato</option>
                <option value='dog'>Perro</option>
              </Select>
                </FormControl>
            <Checkbox
              id="rememberMe"
              name="rememberMe"
              onChange={formik.handleChange}
              isChecked={formik.values.rememberMe}
              colorScheme="purple"
            >
              Recordarme
            </Checkbox>
            <Button type="submit" colorScheme="purple" width="full">
              Registrarse
            </Button>
          </VStack>
        </form>
      </Box>
    </Flex>
          </ModalBody>
          <ModalFooter>
            <Button onClick={onClose}>Volver</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
  )
}

export default RegisterModal