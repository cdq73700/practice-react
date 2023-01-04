import { createApi } from '@reduxjs/toolkit/query/react'
import { staggeredBaseQueryWithBailOut } from '../../../Api'

export const usersApi = createApi({
	reducerPath: 'usersApi',
	baseQuery: staggeredBaseQueryWithBailOut,
	endpoints: (builder) => ({
		userList: builder.query({
			query: () => `users`,
		}),
		userDetail: builder.query({
			query: ({ id }) => `users/${id}`,
		}),
	}),
})

export const { useUserListQuery, useUserDetailQuery } = usersApi
