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
				Id: response.items.Id,
				Email: response.items.Email,
				Name: response.items.Name,
			}),
		}),
		userAdd: builder.mutation({
			query: (body) => ({
				url: 'users/Add',
				method: 'POST',
				body,
			}),
		}),
		userUpdate: builder.mutation({
			query: (body) => ({
				url: 'users/update',
				method: 'PUT',
				body,
			}),
		}),
		userDelete: builder.mutation({
			query: (body) => ({
				url: 'users/delete',
				method: 'DELETE',
				body,
			}),
		}),
	}),
})

export const {
	useUserListQuery,
	useUserDetailQuery,
	usePrefetch,
	useUserAddMutation,
	useUserUpdateMutation,
	useUserDeleteMutation,
} = usersApi
