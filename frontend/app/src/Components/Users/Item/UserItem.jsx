import {
	Button,
	ButtonGroup,
	Card,
	CardBody,
	CardFooter,
	CardHeader,
	Heading,
	HStack,
	Spacer,
	Stack,
	useDisclosure,
} from '@chakra-ui/react'
import { AnimatePresence } from 'framer-motion'
import PropTypes from 'prop-types'
import React, { useCallback, useRef } from 'react'
import { AnimationButtonFade } from '../../Button/Button'
import { Edit } from './Edit'
import { View } from './View'

export const UserItem = ({ Id, Email, Name }) => {
	const { isOpen, onOpen, onClose } = useDisclosure()
	const EditOnClick = useCallback(() => {
		if (!isOpen) {
			onOpen()
		} else {
			onClose()
		}
	}, [isOpen])

	return (
		<Stack p="5px 15px 0px 15px" maxH="100vh">
			<Card variant="outline">
				<CardHeader>
					<HStack>
						<Heading>{Name}</Heading>
						<Spacer />
						<ButtonGroup>
							<Button colorScheme="green" onClick={EditOnClick}>
								Edit
							</Button>
						</ButtonGroup>
					</HStack>
				</CardHeader>
				<CardBody>
					{isOpen ? (
						<Edit Email={Email} Name={Name} />
					) : (
						<View Email={Email} Name={Name} />
					)}
				</CardBody>
				<CardFooter>
					<Spacer />
					<AnimatePresence>
						{isOpen && (
							<ButtonGroup>
								<AnimationButtonFade>Save</AnimationButtonFade>
								<AnimationButtonFade colorScheme="red">
									Delete
								</AnimationButtonFade>
							</ButtonGroup>
						)}
					</AnimatePresence>
				</CardFooter>
			</Card>
		</Stack>
	)
}

UserItem.propTypes = {
	Id: PropTypes.any,
	Email: PropTypes.string,
	Name: PropTypes.string,
}
