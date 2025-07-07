import React from 'react'
import { useNavigate } from 'react-router-dom'

interface Item {
  id: number
  address: string
  title: string
  html_version: string
  internal_links: number
  external_links: number
  broken_links: number
  status: string
}

interface Props {
  items: Item[]
  selected: number[]
  onSelect: (id: number, checked: boolean) => void
  onSort: (key: keyof Item) => void
  sortKey: keyof Item
  sortDir: 'asc' | 'desc'
}

function URLTable({ items, selected, onSelect, onSort, sortKey, sortDir }: Props) {
  const navigate = useNavigate()
  return (
    <table className="results-table">
      <thead>
        <tr>
          <th></th>
          <th onClick={() => onSort('title')}>Title {sortKey==='title'? (sortDir==='asc'?'▲':'▼'):''}</th>
          <th onClick={() => onSort('html_version')}>HTML {sortKey==='html_version'? (sortDir==='asc'?'▲':'▼'):''}</th>
          <th onClick={() => onSort('internal_links')}>Internal {sortKey==='internal_links'? (sortDir==='asc'?'▲':'▼'):''}</th>
          <th onClick={() => onSort('external_links')}>External {sortKey==='external_links'? (sortDir==='asc'?'▲':'▼'):''}</th>
          <th onClick={() => onSort('broken_links')}>Broken {sortKey==='broken_links'? (sortDir==='asc'?'▲':'▼'):''}</th>
          <th onClick={() => onSort('status')}>Status {sortKey==='status'? (sortDir==='asc'?'▲':'▼'):''}</th>
        </tr>
      </thead>
      <tbody>
        {items.map(it => (
          <tr
            key={it.id}
            className={selected.includes(it.id) ? "selected" : ""}
            style={{ cursor: "pointer" }}>
            <td><input type="checkbox" checked={selected.includes(it.id)} onChange={e=>onSelect(it.id,e.target.checked)} /></td>
            <td onClick={() => navigate(`/detail/${it.id}`)}>{it.title}</td>
            <td>{it.html_version}</td>
            <td>{it.internal_links}</td>
            <td>{it.external_links}</td>
            <td>{it.broken_links}</td>
            <td>{it.status}</td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default URLTable

