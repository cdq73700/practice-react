import { Grid, GridItem } from '@chakra-ui/react'
import { useUserListQuery } from '../Api'
import { UserItem } from '../Item/UserItem'
import { UserList } from '../List/UserList'

export const UsersView = () => {
	const userList = useUserListQuery()
	const testData = Array.from({ length: 100 }).map((_, index) => ({
		Id: index,
		Email: `test${index}@test.com`,
		Name: `test${index}`,
		Password: `test${index}`,
	}))
	return (
		<Grid gridTemplateColumns="0.35fr 1fr">
			<GridItem>
				<UserList {...userList} data={{ items: testData }} />
			</GridItem>
			<GridItem>
				<UserItem {...testData[0]} />
			</GridItem>
		</Grid>
	)
}
