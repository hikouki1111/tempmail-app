let currentAccount = "";

async function init() {
    await updateAccounts();
    document.addEventListener("keydown", (event) => {
        if (event.key === "Escape") {
            const mailBodyElement = document.getElementById("mail-body");
            mailBodyElement.remove();
        }
    });
}

async function updateAccounts() {
    const emailList = document.getElementById("email-list");
    emailList.innerHTML = "";
    const accounts = await getAccounts();
    accounts.forEach(acc => {
        addItem(acc.email, acc.token);
    });
}

function addItem(email, token) {
    const emailList = document.getElementById("email-list");
    const emailItem = document.createElement("div");
    emailItem.classList.add("email-item");

    const emailText = document.createElement("p");
    emailText.innerText = email;
    emailText.addEventListener("click", () => {
        currentAccount = emailText.innerText;
        document.getElementById("current-account").innerText = currentAccount;

        const refreshSvg = document.createElement("img");
        refreshSvg.id = "refresh";
        refreshSvg.classList.add("clickable-icon");
        refreshSvg.src = "assets/refresh.svg";
        refreshSvg.style.top = "0";
        refreshSvg.style.right = "0";
        refreshSvg.style.position = "absolute";
        refreshSvg.addEventListener("click", () => {
            updateMailbox();
        });
        if (!document.getElementById("refresh")) {
            document.body.appendChild(refreshSvg);
        }

        updateMailbox();
    });

    const copySvg = document.createElement("img");
    copySvg.classList.add("clickable-icon");
    copySvg.id = email;
    copySvg.src = "assets/copy.svg";
    copySvg.addEventListener("click", () => {
        navigator.clipboard.writeText(copySvg.id);
    });
    copySvg.alt = "copy";

    const deleteSvg = document.createElement("img");
    deleteSvg.classList.add("clickable-icon");
    deleteSvg.id = token;
    deleteSvg.alt = email;
    deleteSvg.src = "assets/delete.svg";
    deleteSvg.addEventListener("click", async () => {
        await deleteAccount(deleteSvg.id);
        await updateAccounts();
        if (deleteSvg.alt === currentAccount) {
            currentAccount = "";
            document.getElementById("refresh").remove();
        }
    });

    emailItem.appendChild(emailText);
    emailItem.appendChild(copySvg);
    emailItem.appendChild(deleteSvg);
    emailList.appendChild(emailItem);
}

async function onClickAddAccount() {
    await addAccount();
    await updateAccounts();
}

async function updateMailbox() {
    const mailbox = await getMailbox(currentAccount);
    const mailboxElement = document.getElementById("mail-box");
    mailboxElement.innerHTML = "";
    mailbox.forEach(mail => {
        const mailJson = JSON.parse(mail);
        const subject = "Subject: " + mailJson.Subject;
        const from = "From " + mailJson.From;

        const mailElement = document.createElement("div");
        mailElement.classList.add("mail");
        mailElement.addEventListener("click", async () => {
            const mailBody =  document.createElement("div");
            mailBody.id = "mail-body";
            mailBody.innerHTML = "";

            const mailBodySubject = document.createElement("p");
            mailBodySubject.innerText = subject;
            mailBody.appendChild(mailBodySubject)
            const mailBodyFrom = document.createElement("p");
            mailBodyFrom.innerText = from;
            mailBody.appendChild(mailBodyFrom);
            if (mailJson.BodyText == "") {
                mailBody.innerHTML += mailJson.BodyHTML;
            } else {
                mailBody.innerHTML += mailJson.BodyText;
            }

            if (mailJson.Attachments.length !== 0) {
                const attachmentsJson = await getAttachments(currentAccount, mailJson.ID);
                const attachments = JSON.parse(attachmentsJson);
                if (Array.isArray(attachments)) {
                    attachments.forEach(a => {
                        const linkElement = document.createElement("a");
                        linkElement.innerText = "Download " + a.Name;
                        linkElement.href = a.URL;
                        mailBody.appendChild(linkElement);
                        mailBody.appendChild(document.createElement("br"));
                    });
                }
            }

            const closeButton = document.createElement("img");
            closeButton.src = "assets/close.svg";
            closeButton.classList.add("clickable-icon");
            closeButton.addEventListener("click", () => {
                const mailBodyElement = document.getElementById("mail-body");
                mailBodyElement.remove();
            });
            closeButton.style.position = "fixed";
            closeButton.style.top = "0";
            closeButton.style.left = "0";

            mailBody.appendChild(closeButton);
            document.body.appendChild(mailBody);
        })
        const mailSubject = document.createElement("p");
        mailSubject.innerText = "Subject: " + mailJson.Subject;
        mailElement.appendChild(mailSubject);

        const mailFrom = document.createElement("p");
        mailFrom.innerText = "From " + mailJson.From;
        mailElement.appendChild(mailFrom);
        mailboxElement.appendChild(mailElement);
    });
}

init();