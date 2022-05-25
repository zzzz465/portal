import type { GetServerSidePropsResult, NextPage } from 'next'
import Image from 'next/image'
import noImage from '../../public/no-image.png'
import { DisplayableRecordItem, GroupedRecordItem, WellKnownRecordItem } from '../types/recordItem'
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

const Home: NextPage<Props> = (props: Props) => {
  if (props.error) {
    // TODO: error handling
  }

  const WellKnownRecordItem = (item: WellKnownRecordItem) => (
    <a className='block' href={`http://${item.data.name}`}>
      {/* Icon */}
      <Image width='128px' height='128px' className='aspect-square' alt='icon' src={noImage.src} />

      {/* title */}
      {/* <Text> creates hydration error. https://stackoverflow.com/questions/71706064/react-18-hydration-failed-because-the-initial-ui-does-not-match-what-was-render */}
      <h3 className='text-center'>{item.name}</h3>
    </a>
  )

  const GroupedRecordItem = (item: GroupedRecordItem) => (
    <a className='block'>
      {/* Icon */}
      <Image width='128px' height='128px' className='aspect-square' alt='icon' src={noImage.src} />

      {/* title */}
      {/* <Text> creates hydration error. https://stackoverflow.com/questions/71706064/react-18-hydration-failed-because-the-initial-ui-does-not-match-what-was-render */}
      <h3 className='text-center'>{item.name}</h3>
    </a>
  )

  const RecordItem = (item: DisplayableRecordItem) => {
    switch (item.type) {
      case 'groupedRecordItem':
        return GroupedRecordItem(item)
      case 'wellKnownRecordItem':
        return WellKnownRecordItem(item)
      default:
        <div><h3>(Error) {item}</h3></div>
    }
  }

  return (
    // root
    <div style={{  /* backgroundImage: `url(${backgroundImage.src})` */ }} className='w-screen h-screen flex flex-col items-center overflow-hidden'>
      <header style={{ backgroundColor: 'green' }} className='w-full flex flex-col items-center h-14'>
        <div style={{ width: '960px', border: '2px solid red', background: 'white' }} className='h-full flex flex-row items-center my-2'>
          <h1 className='w-full text-center'>
            Search
          </h1>
        </div>
      </header>

      {/* main */}
      <div id='main-container' className='w-full flex flex-row flex-1 overflow-auto my-12 align-middle justify-center'>
        {/* upper layout */}
        {/* <div id='upper-layout' style={{}} className='w-full h-2/5 flex flex-row'> */}
        {/* frequently accessed records */}
        {/* TODO: prevent growing */}
        {/* <div id='frequently-accessed-records' style={{ flex: '1' }} className='p-2 px-4 overflow-y-scroll'>
            <ol>
              {items()}
            </ol>
          </div> */}

        {/* recntly added records */}
        {/* <div id='recently-added-records' style={{ flex: '1' }} className='p-2 px-4 overflow-y-scroll'>
            <ol>
              {items()}
            </ol>
          </div> */}
        {/* </div> */}

        {/* bottom layout */}
        <div id='bottom-layout' style={{ flex: 1 }} className='w-3/4'>
          <div id='grid-items-container' className='mx-auto p-8'>
            {/* all groups (grid icon list) */}
            <ol className='w-full flex flex-row flex-wrap gap-12 overflow-y-hidden justify-start'>
              {
                props.items.map((item, i) =>
                  <li key={i}>
                    {RecordItem(item)}
                  </li>
                )
              }
            </ol>
          </div>
        </div>
      </div>
    </div >
  )
}

export default Home
