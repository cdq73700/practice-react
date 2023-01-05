import { useContext } from 'react'
import { UserContext, UserDispatch } from './Reducers'

export const useSelectUser = () => {
	const state = useContext(UserContext)
	const dispatch = useContext(UserDispatch)

	const updateUser = (id) => {
		dispatch({ type: 'SELECT_USER', payload: id })
	}

	return { userId: state.userId, updateUser }
}
