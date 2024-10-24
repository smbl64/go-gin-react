import { useEffect, useState } from 'react'
import './App.css'

import useSWR from 'swr';

const srvAddr = import.meta.env.VITE_SERVER_ADDR;
const fetcher = (url) => fetch(`${srvAddr}/${url}`).then((res) => res.json());

function App() {


  const { data: user, isLoading, error } = useSWR('api/user', fetcher);

  if (error) return <div>failed to load</div>
  if (isLoading) return <div>loading...</div>

  return (
    <>
      <p> Server URL is {srvAddr} </p>
      <div>
        <p>Username: {user.username}</p>
        <p>Email: {user.email}</p>
      </div>

    </>
  )
}

export default App
