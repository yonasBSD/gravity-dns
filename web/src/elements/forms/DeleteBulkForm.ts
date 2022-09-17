import { CSSResult, TemplateResult, html } from "lit";
import { customElement, property, state } from "lit/decorators.js";
import { until } from "lit/directives/until.js";

import PFList from "@patternfly/patternfly/components/List/list.css";

import { EVENT_REFRESH } from "../../common/constants";
import { MessageLevel } from "../../common/messages";
import { PFSize } from "../Spinner";
import { ModalButton } from "../buttons/ModalButton";
import "../buttons/SpinnerButton";
import { showMessage } from "../messages/MessageContainer";
import { PaginatedResponse } from "../table/Table";
import { Table, TableColumn } from "../table/Table";

type BulkDeleteMetadata = { key: string; value: string }[];

@customElement("ak-forms-delete-bulk")
export class DeleteBulkForm extends ModalButton {
    @property({ attribute: false })
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    objects: any[] = [];

    @property()
    objectLabel?: string;

    @property({ attribute: false })
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    metadata: (item: any) => BulkDeleteMetadata = (item: any) => {
        const rec = item as Record<string, unknown>;
        const meta = [];
        if (Object.prototype.hasOwnProperty.call(rec, "name")) {
            meta.push({ key: `Name`, value: rec.name as string });
        }
        if (Object.prototype.hasOwnProperty.call(rec, "pk")) {
            meta.push({ key: `ID`, value: rec.pk as string });
        }
        return meta;
    };

    @property({ attribute: false })
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    delete!: (item: any) => Promise<any>;

    confirm(): Promise<void> {
        return Promise.all(
            this.objects.map((item) => {
                return this.delete(item);
            }),
        )
            .then(() => {
                this.onSuccess();
                this.open = false;
                this.dispatchEvent(
                    new CustomEvent(EVENT_REFRESH, {
                        bubbles: true,
                        composed: true,
                    }),
                );
            })
            .catch((e) => {
                this.onError(e);
                throw e;
            });
    }

    onSuccess(): void {
        showMessage({
            message: `Successfully deleted ${this.objects.length} ${this.objectLabel}`,
            level: MessageLevel.success,
        });
    }

    onError(e: Error): void {
        showMessage({
            message: `Failed to delete ${this.objectLabel}: ${e.toString()}`,
            level: MessageLevel.error,
        });
    }

    renderModalInner(): TemplateResult {
        return html`<section class="pf-c-modal-box__header pf-c-page__main-section pf-m-light">
                <div class="pf-c-content">
                    <h1 class="pf-c-title pf-m-2xl">${`Delete ${this.objectLabel}`}</h1>
                </div>
            </section>
            <section class="pf-c-modal-box__body pf-c-page__main-section pf-m-light">
                <form class="pf-c-form pf-m-horizontal">
                    <p class="pf-c-title">
                        ${`Are you sure you want to delete ${this.objects.length} ${this.objectLabel}?`}
                    </p>
                    <slot name="notice"></slot>
                </form>
            </section>
            <footer class="pf-c-modal-box__footer">
                <ak-spinner-button
                    .callAction=${() => {
                        return this.confirm();
                    }}
                    class="pf-m-danger"
                >
                    ${`Delete`} </ak-spinner-button
                >&nbsp;
                <ak-spinner-button
                    .callAction=${async () => {
                        this.open = false;
                    }}
                    class="pf-m-secondary"
                >
                    ${`Cancel`}
                </ak-spinner-button>
            </footer>`;
    }
}
