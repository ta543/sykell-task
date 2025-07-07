import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import axios from 'axios'

interface URLItem {
  id: number
  address: string
  title: string
  html_version: string
  internal_links: number
  external_links: number
  broken_links: number
  broken_links_detail: { link: string; status_code: number }[]
  status: string
  h1_count: number
  h2_count: number
  h3_count: number
  has_login_form: boolean
}

function Detail() {
  const { id } = useParams()
  const [data, setData] = useState<URLItem | null>(null)

  useEffect(() => {
    axios
      .get<URLItem>(`/api/urls/${id}`, { headers: { Authorization: 'secret-token' } })
      .then(res => setData(res.data))
      .catch(err => console.error(err))
  }, [id])

  if (!data) return <div>Loading...</div>
  return (
    <div className="detail-container">
      <h2>{data.title}</h2>
      <p>URL: {data.address}</p>
      <p>HTML Version: {data.html_version}</p>
      <p>H1: {data.h1_count}, H2: {data.h2_count}, H3: {data.h3_count}</p>
      <div style={{ width: 200, height: 200, position: 'relative' }}>
        <div style={{ position: 'absolute', width: '100%', height: '100%', borderRadius: '50%', background: '#eee' }}></div>
        <div
          style={{
            position: 'absolute',
            width: '100%',
            height: '100%',
            borderRadius: '50%',
            clip: 'rect(0, 200px, 200px, 100px)',
            background: '#4caf50',
            transform: `rotate(${(data.internal_links / (data.internal_links + data.external_links)) * 360}deg)`
          }}
        ></div>
      </div>
      <p>Internal Links: {data.internal_links}</p>
      <p>External Links: {data.external_links}</p>
      <p>Broken Links: {data.broken_links}</p>
      {data.broken_links_detail && data.broken_links_detail.length > 0 && (
        <ul>
          {data.broken_links_detail.map((b, idx) => (
            <li key={idx}>
              {b.link} - {b.status_code}
            </li>
          ))}
        </ul>
      )}
      <p>Login Form: {data.has_login_form ? 'Yes' : 'No'}</p>
      <p>Status: {data.status}</p>
    </div>
  )
}

export default Detail
