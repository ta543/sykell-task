import React, { useEffect, useState } from 'react'
import axios from 'axios'
import URLTable from '../components/URLTable'

interface URLItem {
  id: number
  address: string
  title: string
  html_version: string
  internal_links: number
  external_links: number
  broken_links: number
  status: string
}

function Dashboard() {
  const [urls, setUrls] = useState<URLItem[]>([])
  const [newUrl, setNewUrl] = useState('')
  const [search, setSearch] = useState('')
  const [page, setPage] = useState(1)
  const [limit, setLimit] = useState(10)
  const [selected, setSelected] = useState<number[]>([])
  const [statusFilter, setStatusFilter] = useState('')
  const [htmlFilter, setHtmlFilter] = useState('')
  const [sortKey, setSortKey] = useState<keyof URLItem>('title')
  const [sortDir, setSortDir] = useState<'asc' | 'desc'>('asc')

  const fetchUrls = async () => {
    try {
      const res = await axios.get<URLItem[]>(
        `/api/urls?page=${page}&limit=${limit}&search=${encodeURIComponent(search)}&status=${statusFilter}&html=${htmlFilter}`,
        { headers: { Authorization: 'secret-token' } }
      )
      let list = res.data
      list = list.sort((a, b) => {
        const aVal = a[sortKey]
        const bVal = b[sortKey]
        if (aVal < bVal) return sortDir === 'asc' ? -1 : 1
        if (aVal > bVal) return sortDir === 'asc' ? 1 : -1
        return 0
      })
      setUrls(list)
    } catch (e) {
      console.error(e)
    }
  }

  useEffect(() => {
    fetchUrls()
    const t = setInterval(fetchUrls, 5000)
    return () => clearInterval(t)
  }, [page, limit, search, statusFilter, htmlFilter, sortKey, sortDir])

  const addUrl = async () => {
    if (!newUrl) return
    await axios.post('/api/urls', { address: newUrl }, { headers: { Authorization: 'secret-token' } })
    setNewUrl('')
    fetchUrls()
  }

  const onSelect = (id: number, checked: boolean) => {
    setSelected(prev => checked ? [...prev, id] : prev.filter(i => i !== id))
  }

  const onSort = (key: keyof URLItem) => {
    if (sortKey === key) {
      setSortDir(d => (d === 'asc' ? 'desc' : 'asc'))
    } else {
      setSortKey(key)
      setSortDir('asc')
    }
  }

  const restartSelected = async () => {
    for (const id of selected) {
      await axios.post(`/api/urls/${id}/restart`, {}, { headers: { Authorization: 'secret-token' } })
    }
    setSelected([])
    fetchUrls()
  }

  const deleteSelected = async () => {
    for (const id of selected) {
      await axios.delete(`/api/urls/${id}`, { headers: { Authorization: 'secret-token' } })
    }
    setSelected([])
    fetchUrls()
  }

  const stopSelected = async () => {
    for (const id of selected) {
      await axios.post(`/api/urls/${id}/stop`, {}, { headers: { Authorization: 'secret-token' } })
    }
    setSelected([])
    fetchUrls()
  }

  return (
    <div>
      <div className="controls">
        <input value={newUrl} onChange={e => setNewUrl(e.target.value)} placeholder="Enter URL" />
        <button onClick={addUrl}>Add</button>
        <button onClick={fetchUrls}>Refresh</button>
      </div>
      <div className="controls">
        <input value={search} onChange={e => setSearch(e.target.value)} placeholder="Search" />
        <select value={statusFilter} onChange={e => setStatusFilter(e.target.value)}>
          <option value="">All Status</option>
          <option value="queued">Queued</option>
          <option value="processing">Processing</option>
          <option value="done">Done</option>
          <option value="error">Error</option>
          <option value="stopped">Stopped</option>
        </select>
        <select value={htmlFilter} onChange={e => setHtmlFilter(e.target.value)}>
          <option value="">All HTML</option>
          <option value="HTML5">HTML5</option>
          <option value="HTML4">HTML4</option>
        </select>
        <select value={limit} onChange={e => setLimit(parseInt(e.target.value))}>
          <option value={5}>5</option>
          <option value={10}>10</option>
          <option value={20}>20</option>
        </select>
        <button onClick={restartSelected} disabled={!selected.length}>Re-run</button>
        <button onClick={stopSelected} disabled={!selected.length}>Stop</button>
        <button onClick={deleteSelected} disabled={!selected.length}>Delete</button>
      </div>
      <div className="table-wrapper">
        <URLTable items={urls} selected={selected} onSelect={onSelect} onSort={onSort} sortKey={sortKey} sortDir={sortDir} />
      </div>
      <div className="controls">
        <button onClick={() => setPage(p => Math.max(1, p - 1))}>Prev</button>
        <span>Page {page}</span>
        <button onClick={() => setPage(p => p + 1)}>Next</button>
      </div>
    </div>
  )
}

export default Dashboard
