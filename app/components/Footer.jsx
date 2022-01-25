import styles from '../styles/Footer.module.css';
import Image from 'next/image'
import Grid from '@mui/material/Grid';


export default function Footer() {
    return (
    <Grid>
        <footer className={styles.footer}>
            <a
                href="https://github.com/Nexlson"
                target="_blank"
                rel="noopener noreferrer"
                >
                <p>Built by Nexlson</p> {' '}
                <span className={styles.logo}>
                    <Image src="/github.ico" alt="Github Logo" width={16} height={16} />
                </span>
            </a>
        </footer>
    </Grid>
    )
}


