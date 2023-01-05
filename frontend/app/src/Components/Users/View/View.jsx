import { Grid, GridItem } from '@chakra-ui/react'
import { skipToken } from '@reduxjs/toolkit/dist/query'
import { useUserDetailQuery, useUserListQuery } from '../Api'
import { UserItem } from '../Item/UserItem'
import { UserList } from '../List/UserList'
import { useSelectUser } from '../Reducers/CustomHook'

export const UsersView = () => {
	const userList = useUserListQuery()
	const { userId } = useSelectUser()
	const userDetail = useUserDetailQuery(userId !== '' ? userId : skipToken)
	const testData = Array.from({ length: 100 }).map((_, index) => ({
		Id: index,
		Email: `test${index}@test.com`,
		Name: `test${index}`,
		Password: `test${index}`,
	}))

	return (
		<Grid gridTemplateColumns="0.35fr 1fr">
			<GridItem>
				<UserList {...userList} />
			</GridItem>
			<GridItem>{userId !== '' && <UserItem {...userDetail} />}</GridItem>
		</Grid>
	)
}
