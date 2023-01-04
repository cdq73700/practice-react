import { configureStore } from '@reduxjs/toolkit'
import { setupListeners } from '@reduxjs/toolkit/query/react'
import { backendApi } from '../Api'
import { usersApi } from '../Components/Users/Api'

export const store = configureStore({
	reducer: {
		// Add the generated reducer as a specific top-level slice
		[backendApi.reducerPath]: backendApi.reducer,
		[usersApi.reducerPath]: usersApi.reducer,
	},
	// Adding the api middleware enables caching, invalidation, polling,
	// and other useful features of `rtk-query`.
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware()
			.concat(backendApi.middleware)
			.concat(usersApi.middleware),
})

// optional, but required for refetchOnFocus/refetchOnReconnect behaviors
// see `setupListeners` docs - takes an optional callback as the 2nd arg for customization
setupListeners(store.dispatch)
