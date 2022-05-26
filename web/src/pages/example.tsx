import { GetServerSidePropsResult } from 'next'
import Image from 'next/image'
import { VscSearch } from 'react-icons/vsc'
import NoImage from '../../public/no-image.png'
import { Displayable, DisplayableRecordItem, GroupedRecordItem, WellKnownRecordItem } from '../types/recordItem'
import { tryP } from '../utils/try'

type Props = {
  items: DisplayableRecordItem[]
  error?: Error
}

export async function getServerSideProps(): Promise<GetServerSidePropsResult<Props>> {
  const [res, err0] = await tryP(() => fetch('http://localhost:3000/api/items'))
  if (err0 !== null) {
    return { props: { items: [], error: err0 as Error } }
  }

  const data = await res.json()
  return { props: data }
}

export default function Example(props: Props) {
  const wellKnownItems = props.items.filter((item) => item.type === 'wellKnownRecordItem') as WellKnownRecordItem[]
  const groupedItems = props.items.filter((item) => item.type === 'groupedRecordItem') as GroupedRecordItem[]
  const items = [...wellKnownItems, ...groupedItems]

  const appItem = (item: Displayable) => (
    <a className="flex flex-col items-center gap-3 hover:shadow-2xl" href={`http://${item.name}`}>
      <Image src={NoImage.src} width="52px" height="52px" alt="no-image" layout="fixed"></Image>
      <h3>{item.name}</h3>
    </a>
  )

  return (
    <div id="base" className="w-screen h-screen">
      <header id="header" className="w-full h-14 bg-slate-700 flex justify-center items-center">
        <div id="search" className="bg-slate-300 w-1/3 h-1/2 flex itmes-center gap-1 pl-1">
          <VscSearch className="my-auto" />
          <input placeholder="Search" className="bg-slate-300 w-full h-full p-1"></input>
        </div>
      </header>

      {/* main layout */}
      <div className="w-2/3 mx-auto py-12">
        {/* known applications */}
        <div id="known-applications" className="flex flex-col gap-8">
          {/* title */}
          <h1 className="text-xl font-bold">Applications</h1>

          {/* icons layout */}
          <div>
            <ol className="flex flex-row gap-6">
              {items.map((item) => (
                <li key={item.name} style={{ width: '128px', height: '96px' }}>
                  {appItem(item)}
                </li>
              ))}
            </ol>
          </div>
        </div>

        {/* application groups */}
      </div>
    </div>
  )
}
