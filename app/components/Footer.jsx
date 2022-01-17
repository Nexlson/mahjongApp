import styles from '../styles/Home.module.css'
import Image from 'next/image'

export default function Footer() {
    return (
    <footer className={styles.footer}>
        <a
            href="https://github.com/Nexlson"
            target="_blank"
            rel="noopener noreferrer"
            >
            Built by Nexlson{' '}
            <span className={styles.logo}>
                <Image src="/github.ico" alt="Github Logo" width={16} height={16} />
            </span>
        </a>
    </footer>
    )
}


