import range from './utils'
import Holder from './Holder'

export default function HolderGroup() {
    const holders = range(1,14)

    return (
        <>
            {holders.map((holder, index) => <Holder id={holder} holdId={index}/>)}
        </>
    )
}