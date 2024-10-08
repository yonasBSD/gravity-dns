import { CSSResult } from "lit";
import { TemplateResult, html } from "lit";

import PFContent from "@patternfly/patternfly/components/Content/content.css";
import PFPage from "@patternfly/patternfly/components/Page/page.css";
import PFSidebar from "@patternfly/patternfly/components/Sidebar/sidebar.css";

import { EVENT_TMP_TITLE } from "../../common/constants";
import { updateURLParams } from "../router/RouteMatch";
import { Table } from "../table/Table";

export abstract class TablePage<T extends object> extends Table<T> {
    abstract pageTitle(): string;
    abstract pageDescription(): string | undefined;
    abstract pageIcon(): string;

    static get styles(): CSSResult[] {
        return super.styles.concat(PFPage, PFContent, PFSidebar);
    }

    renderSidebarBefore(): TemplateResult {
        return html``;
    }

    renderSidebarAfter(): TemplateResult {
        return html``;
    }

    renderEmpty(inner?: TemplateResult): TemplateResult {
        return super.renderEmpty(html`
            ${inner
                ? inner
                : html`<ak-empty-state icon=${this.pageIcon()} header="${"No objects found."}">
                      <div slot="body">
                          ${this.searchEnabled() ? this.renderEmptyClearSearch() : html``}
                      </div>
                      <div slot="primary">${this.renderObjectCreate()}</div>
                  </ak-empty-state>`}
        `);
    }

    renderEmptyClearSearch(): TemplateResult {
        if (this.search === "") {
            return html``;
        }
        return html`<button
            @click=${() => {
                this.search = "";
                this.requestUpdate();
                this.fetch();
                updateURLParams({
                    search: "",
                });
            }}
            class="pf-v6-c-button pf-m-link"
        >
            ${"Clear search"}
        </button>`;
    }

    renderObjectCreate(): TemplateResult {
        return html``;
    }

    renderToolbar(): TemplateResult {
        return html`${this.renderObjectCreate()}${super.renderToolbar()}`;
    }

    render(): TemplateResult {
        this.dispatchEvent(
            new CustomEvent(EVENT_TMP_TITLE, {
                bubbles: true,
                composed: true,
                detail: {
                    title: this.pageTitle(),
                    icon: this.pageIcon(),
                    subtext: this.pageDescription(),
                },
            }),
        );
        return html`<section class="pf-v6-c-page__main-section pf-m-no-padding-mobile">
            <div class="pf-v6-c-sidebar pf-m-gutter">
                <div class="pf-v6-c-sidebar__main">
                    ${this.renderSidebarBefore()}
                    <div class="pf-v6-c-sidebar__content">
                        <div class="pf-v6-c-card">
                            <div class="pf-v6-c-card__body">${this.renderTable()}</div>
                        </div>
                    </div>
                    ${this.renderSidebarAfter()}
                </div>
            </div>
        </section>`;
    }
}
