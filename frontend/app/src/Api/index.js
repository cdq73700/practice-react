import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

export const usersApi = createApi({
	reducerPath: 'usersApi',
	baseQuery: fetchBaseQuery({ baseUrl: 'http://localhost:4000/api/v1/' }),
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
