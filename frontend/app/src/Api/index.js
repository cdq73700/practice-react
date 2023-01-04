import { createApi, fetchBaseQuery, retry } from '@reduxjs/toolkit/query/react'

export const staggeredBaseQueryWithBailOut = retry(
	async (args, api, extraOptions) => {
		const result = await fetchBaseQuery({
			baseUrl: 'http://localhost:4000/api/v1/',
		})(args, api, extraOptions)

		// bail out of re-tries immediately if unauthorized,
		// because we know successive re-retries would be redundant
		if (
			['FETCH_ERROR', 'PARSING_ERROR', 'TIMEOUT_ERROR'].some(
				(item) => item === result.error?.status
			)
		) {
			retry.fail(result.error)
		}

		return result
	},
	{
		maxRetries: 5,
	}
)

export const backendApi = createApi({
	reducerPath: 'backendApi',
	baseQuery: staggeredBaseQueryWithBailOut,
	endpoints: () => ({}),
})
