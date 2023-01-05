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
			query: (id) => `users/${id}`,
			transformResponse: (response) => ({
				Id: response.items[0].Id,
				Email: response.items[0].Email,
				Name: response.items[0].Name,
			}),
		}),
	}),
})

export const { useUserListQuery, useUserDetailQuery, usePrefetch } = usersApi